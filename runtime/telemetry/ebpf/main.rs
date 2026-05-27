use aya::Bpf;

fn main() {
    let mut bpf = Bpf::load_file("cfos-ebpf.o").unwrap();

    println!("eBPF telemetry active");
}