# Test

harbour -gc -I/usr/local/include/harbour linkaj_me.prg
harbour -gh hrbext.prg hrbext2.prg
harbour -gh pg_test.prg
harbour -I/usr/local/include/harbour -gh pg_object_test.prg
harbour -gh run_main.prg
go test
