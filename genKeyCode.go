package main

func genKeyCode() string {
	//should be 94, will be 100 so that you have a little skim on the top as a buffer
	return genRandomString(len(clear) - 1)
}
