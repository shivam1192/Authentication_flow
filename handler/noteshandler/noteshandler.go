package noteshandler

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"../../model/notesmodel"
	"../../databaseConnect"
)


func GetAllNotes(w http.ResponseWriter, r *http.Request){
	c := r.Context().Value("claim")
	var notes []notesmodel.Notes
	databaseConnect.DB.Find(&notes,"Users_ID=?", int(c.(uint)));
	json.NewEncoder(w).Encode(&notes);
}

func CreateNote(w http.ResponseWriter, r *http.Request){
	c := r.Context().Value("claim")
	var note notesmodel.Notes;
	json.NewDecoder(r.Body).Decode(&note);
	note.UsersID = int(c.(uint))
	createdNote := databaseConnect.DB.Create(&note);
	err := createdNote.Error
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}else{
		json.NewEncoder(w).Encode(&note)
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request){
	c := r.Context().Value("claim")
	params := mux.Vars(r);
	var note notesmodel.Notes
	databaseConnect.DB.First(&note, params["id"]);
	if note.UsersID != int(c.(uint)) || note.ID == 0{ 
		w.WriteHeader(http.StatusNotFound)
	}else{
			databaseConnect.DB.Delete(&note);
			json.NewEncoder(w).Encode(&note)	
	}
}

