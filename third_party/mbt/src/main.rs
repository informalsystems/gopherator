use std::collections::BTreeMap;

fn main() {
    let args = std::env::args().skip(1).take(2).collect::<Vec<_>>();
    let tla_tests_file_path = &args[0];
    let tla_config_file_path = &args[1];
    let runtime = modelator::ModelatorRuntime::default();
    let trace_results = runtime
        .traces(tla_tests_file_path, tla_config_file_path)
        .unwrap();
    let trace_results = trace_results
        .into_iter()
        .map(|(testname, traces)| {
            let traces = traces.unwrap();
            (
                testname,
                traces
                    .into_iter()
                    .map(|trace| trace.into_iter().collect::<Vec<_>>())
                    .collect::<Vec<_>>(),
            )
        })
        .collect::<BTreeMap<_, _>>();
    println!("{}", serde_json::to_string_pretty(&trace_results).unwrap());
}
