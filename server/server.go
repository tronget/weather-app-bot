package server

//http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
//
//	switch r.Method {
//	case http.MethodGet:
//		jsn, err := json.Marshal(data)
//		if err != nil {
//			msg := "Internal server error: unable to marshal json"
//			log.Println(msg)
//			http.Error(w, msg, http.StatusInternalServerError)
//			return
//		}
//		w.Write(jsn)
//	case http.MethodPost:
//		bodyBytes, err := io.ReadAll(r.Body)
//		if err != nil {
//			log.Println(err)
//			http.Error(w, "Error reading request body", http.StatusBadRequest)
//		}
//		defer r.Body.Close()
//
//		rd := RequestData{}
//		if err := json.Unmarshal(bodyBytes, &rd); err != nil {
//
//		}
//		log.Println(rd)
//	default:
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//	}
//})
//log.Fatal(http.ListenAndServe(":8080", nil))
