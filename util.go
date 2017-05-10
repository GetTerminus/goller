package goller

func checkErr(e error, l Logger) {
	if e != nil {
		l.Fatal(e)
	}
}
