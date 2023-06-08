use git2::{Repository, ErrorCode};
use std::env;
use std::path::PathBuf;

fn main() -> Result<(), Box<dyn std::error::Error>> {
   let out_dir = PathBuf::from(env::var("OUT_DIR").unwrap());
   let url = "https://github.com/protocolbuffers/protobuf.git";
   match Repository::clone(url, "./protobuf") {
       Ok(_repo) => {},
       Err(e) => {
           if e.code() != ErrorCode::Exists {
               panic!("failed to clone: {}", e)
           }
       }
   };

   tonic_build::configure()
           .protoc_arg("--experimental_allow_proto3_optional") // for older systems
           .build_client(true)
           .build_server(true)
           .file_descriptor_set_path(out_dir.join("store_descriptor.bin"))
           .out_dir("./src")
           .compile(&["./nodeapi.proto"], &["./", "./protobuf/src/"])?;

   Ok(())
}
