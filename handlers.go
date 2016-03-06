package main

type Profile struct {
	Name    string
	Hobbies []string
}

// func Positive(w http.ResponseWriter, r *http.Request) {
// 	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
//
// 	js, err := json.Marshal(profile)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
// }

// func Negative(w http.ResponseWriter, r *http.Request) {
// 	profile := Profile{"Ricky", []string{"cats", "cats"}}
//
// 	js, err := json.Marshal(profile)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
// }
