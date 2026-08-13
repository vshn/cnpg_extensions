metadata = {
  name                     = "plpython3u"
  sql_name                 = "plpython3u"
  image_name               = "plpython3u"
  licenses                 = [ "BSD-2-clause", "BSD-3-Clause", "Custom-Unicode",
                               "Custom-pg_dump", "Custom-regex", "PostgreSQL",
                               "Tcl", "double-metaphone", "nagaysau-ishii" ]
  shared_preload_libraries = []
  extension_control_path   = []
  dynamic_library_path     = []
  ld_library_path          = ["system"]
  auto_update_os_libs      = true
  required_extensions      = []
  create_extension         = true

  versions = {
    bookworm = {
        // renovate: suite=bookworm-pgdg depName=postgresql-plpython3-18
        "18" = "18.6-1.pgdg12+2"
    }
    trixie = {
        // renovate: suite=trixie-pgdg depName=postgresql-plpython3-18
        "18" = "18.6-1.pgdg13+2"
    }
  }
}
