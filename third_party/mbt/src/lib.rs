use std::collections::BTreeMap;
use std::ffi::CString;
use std::os::raw::c_char;

pub fn generate_json_traces_from_tla_tests(
    tla_tests_file_path: &str,
    tla_config_file_path: &str,
) -> Result<String, anyhow::Error> {
    let runtime = modelator::ModelatorRuntime::default();
    let trace_results = runtime.traces(tla_tests_file_path, tla_config_file_path)?;
    let trace_results = trace_results
        .into_iter()
        .map::<Result<_, anyhow::Error>, _>(|(testname, traces)| {
            let traces = traces?;
            Ok((
                testname,
                traces
                    .into_iter()
                    .map(|trace| trace.into_iter().collect::<Vec<_>>())
                    .collect::<Vec<_>>(),
            ))
        })
        .collect::<Result<BTreeMap<_, _>, _>>()?;
    Ok(serde_json::to_string_pretty(&trace_results)?)
}

#[repr(C)]
pub struct CResult {
    json: *mut c_char,
    error: *mut c_char,
}

#[no_mangle]
pub extern "C" fn generate_json_traces_from_tla_tests_rs(
    tla_tests_file_path_c: *mut c_char,
    tla_config_file_path_c: *mut c_char,
) -> CResult {
    let tla_tests_file_path = unsafe { CString::from_raw(tla_tests_file_path_c) };
    let tla_config_file_path = unsafe { CString::from_raw(tla_config_file_path_c) };
    let (json_string, error_string) = match generate_json_traces_from_tla_tests(
        tla_tests_file_path.to_str().unwrap(),
        tla_config_file_path.to_str().unwrap(),
    ) {
        Ok(json) => (json, String::new()),
        Err(error) => (String::new(), format!("{:?}", error)),
    };
    CResult {
        json: CString::new(json_string).unwrap().into_raw(),
        error: CString::new(error_string).unwrap().into_raw(),
    }
}
