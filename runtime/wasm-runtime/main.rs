use wasmtime::*;

fn main() {

    let engine = Engine::default();

    let module = Module::from_file(
        &engine,
        "workload.wasm",
    ).unwrap();

    let mut store = Store::new(&engine, ());

    let instance = Instance::new(
        &mut store,
        &module,
        &[],
    ).unwrap();

    let run = instance
        .get_typed_func::<(), ()>(
            &mut store,
            "run",
        )
        .unwrap();

    run.call(&mut store, ()).unwrap();
}