package goller

func checkErr(e error, l *CustomLogger) {
	if e != nil {
		l.Logger.Fatal(e)
	}
}
