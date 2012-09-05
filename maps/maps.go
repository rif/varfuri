package maps

import (
	//"fmt"
	"appengine"
	"appengine/mail"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var (
	maps = []string{"avramiancu_bradatel.jpg",
		"avramiancu_dealu_livada_deal.jpg",
		"avramiancu_dulalau.jpg",
		"avramiancu_generala.jpg",
		"avramiancu_glemeia.jpg",
		"avramiancu_index_1.jpg",
		"avramiancu_index_2.jpg",
		"avramiancu_index_3.jpg",
		"avramiancu_index_4.jpg",
		"avramiancu_index_5.jpg",
		"avramiancu_index_6.jpg",
		"avramiancu_index_7.jpg",
		"avramiancu_index_8.jpg",
		"avramiancu_index_9.jpg",
		"avramiancu_lunca.jpg",
		"avramiancu_magura_tacasele.jpg",
		"avramiancu_negulesti.jpg",
		"avramiancu_tacasele.jpg",
		"varfurile_casoaie_1.jpg",
		"varfurile_casoaie_2.jpg",
		"varfurile_casoaie_3.jpg",
		"varfurile_glemeie_1.jpg",
		"varfurile_glemeie_2.jpg",
		"varfurile_heread.jpg",
		"varfurile_heredut.jpg",
		"varfurile_index_10.jpg",
		"varfurile_index_11.jpg",
		"varfurile_index_12.jpg",
		"varfurile_index_13.jpg",
		"varfurile_index_14.jpg",
		"varfurile_index_1.jpg",
		"varfurile_index_2.jpg",
		"varfurile_index_3.jpg",
		"varfurile_index_4.jpg",
		"varfurile_index_5.jpg",
		"varfurile_index_6.jpg",
		"varfurile_index_7.jpg",
		"varfurile_index_8.jpg",
		"varfurile_index_9.jpg",
		"varfurile_la_cris.jpg",
		"varfurile_magura_1.jpg",
		"varfurile_magura_2.jpg",
		"varfurile_mlaca.jpg",
		"varfurile_zaloage.jpg",
		"vidra_generala.jpg",
		"vidra_ilicesti.jpg",
		"vidra_index_1.jpg",
		"vidra_index_2.jpg",
		"vidra_index_3.jpg",
		"vidra_index_4.jpg",
		"vidra_index_5.jpg",
		"vidra_index_6.jpg",
		"vidra_index_7.jpg",
		"vidra_index_8.jpg",
		"vidra_mageresti_mihutesti.jpg",
		"vidra_mageresti.jpg",
		"vidra_mihutesti.jpg",
		"vidra_petrisoresti.jpg",
		"vidra_ceva.jpg"}
)

func PrettyName(fn string)(pretty string) {
	parts := strings.Split(fn , "_")
	for _, p := range(parts){
		pretty += strings.TrimRight(p, ".jpg") + " "
	}
	return
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("maps").Funcs(template.FuncMap{"pretty": PrettyName}).ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
	err = t.ExecuteTemplate(w, "base.html", map[string][]string{"Maps":maps})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hartaPage(w http.ResponseWriter, r *http.Request) {
	harta := r.FormValue("h") 
	t, err := template.ParseFiles("templates/base.html", "templates/harta.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, harta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if r.Method == "POST" {
		name := r.FormValue("from")
		email := r.FormValue("email")
		content := r.FormValue("content")
		msg := &mail.Message{
			Sender:  "Radu Fericean (Varfurile Harti) <fericean@gmail.com>",
			To:      []string{"Radu Fericean <radu@fericean.ro>"},
			Subject: fmt.Sprintf("Mesaj Varfuri0Harti de la %s (%s)", name, email),
			Body:    content,
		}
		if err := mail.Send(c, msg); err != nil {
			c.Errorf("Couldn't send email: %v", err)
		}
		return
	}
	t, _ := template.ParseFiles("templates/base.html", "templates/contact.html")
	t.Execute(w, nil)
}

func init() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/harta", hartaPage)
	http.HandleFunc("/contact", contactPage)
}
