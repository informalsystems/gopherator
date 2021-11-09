typedef struct CResult {
  char *json;
  char *error;
} CResult;

struct CResult generate_json_traces_from_tla_tests_rs(char *tla_tests_file_path_c,
                                                      char *tla_config_file_path_c);
