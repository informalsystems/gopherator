use std::collections::BTreeMap;
use std::ffi::CString;
use std::os::raw::c_char;

pub fn generate_json_traces_from_tla_tests(
    tla_tests_file_path: &str,
    tla_config_file_path: &str,
) -> String {
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
    serde_json::to_string_pretty(&trace_results).unwrap()
}

#[no_mangle]
pub extern "C" fn generate_json_traces_from_tla_tests_rs(
    tla_tests_file_path_c: *mut c_char,
    tla_config_file_path_c: *mut c_char,
) -> *mut c_char {
    let tla_tests_file_path = unsafe { CString::from_raw(tla_tests_file_path_c) };
    let tla_config_file_path = unsafe { CString::from_raw(tla_config_file_path_c) };
    let json_string = generate_json_traces_from_tla_tests(
        tla_tests_file_path.to_str().unwrap(),
        tla_config_file_path.to_str().unwrap(),
    );
    CString::new(json_string).unwrap().into_raw()
}
