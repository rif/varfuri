import urllib2

maps = (
		"avramiancu_bradatel.jpg",
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
		"vidra_ceva.jpg",
)

for img in maps:
	print("Downloading: " + img)
	imgurl = 'http://1.varfuri.appspot.com/img/maps/' + img
	req = urllib2.Request(imgurl)
	f = urllib2.urlopen(req)
	img_data = f.read()

	open(img, 'wb').write(img_data)
