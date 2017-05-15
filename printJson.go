func printJson(myInterface interface{}){
	j,err := json.MarshalIndent(myInterface, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(j))
}
