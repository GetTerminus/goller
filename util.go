package goller

func checkErr(e error, l *CustomLogger) {
	if e != nil {
		l.Fatal(e)
	}
}
