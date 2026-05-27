use aya::Bpf;

fn main() {

    let _bpf = Bpf::load_file(
        "cfos-ebpf.o"
    ).unwrap();

    println!("Telemetry active");
}