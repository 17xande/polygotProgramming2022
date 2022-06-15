use projector_rs::opts::Opts;
use clap::Parser;

fn main() {
    let opts = Opts::parse();
    println!("{:?}", opts)
}
