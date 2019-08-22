package smsify

import "strings"

var mapping = [][2]string{ // FIXME: use a .csv file and go generate
	{"salut", "slt"},
	{"ça", "sa"},
	{"coucou", "cc"},
	{"comment", "cmt"},
	{" ?", "?"},
	{"fallait", "falé"},
	{"que", "ke"}, // FIXME: we should only replace the complete word "que", not everytime we meet "que" somewhere in a text
	{"je te", "j'te"},
	{"petite", "ptite"},
	{"petit", "pti"},
	{"que", "ke"},
	{"je t'aime", "jtaime"},
	{"quoi", "koi"},
	{"venir", "vnir"},
	{"quand", "kan"},
	{"unique", "unik"},
	{"d'aussi", "doci"},
	{"aussi", "oci"},
	{"tout", "tt"},
	{"devant", "dvan"},
	{"t'as", "ta"},
	{"j'ai", "jé"},
	{"dehors", "dehor"},
	{"je vais", "jv"},
	{"pour", "pr"},
	{"parceque", "pck"},
	{"je viens", "jv1"},
	{"d'autre", "dotr"},
	{"autre", "otr"},
	{"t'inquiete", "tkt"},
	{"peut-être", "ptet"},
	{"moins", "moin"},
	{"vrai", "vré"},
	{"jamais", "jamé"},
	{"connais", "coné"},
	{"connait", "coné"},
	{"bisous", "bisou"},
	{"mon", "mn"},
	{"l'heure", "lheur"},
	{"gros", "gro"},
	{"qui", "ki"},
	{"c'était", "cété"},
	{"claire", "klr"},
	{"c'est", "c"},
	{"plus", "plu"},
	{"t'es", "t"},
	{"je voulais", "jvoulé"},
	{"voulais", "voulé"},
	{"voulait", "voulé"},
	{"comme", "kom"},
	{"commentaire", "kom"},
	{"quitter", "kité"},
	{"qu'on", "kn"},
	{"quelqu'un", "qqn"},
	{"j'ai", "g"},
	{"qu'a", "ka"},
	{"demain", "dem1"},
	{"manque", "mank"},
	{"aller", "allé"},
	{"cinéma", "ciné"},
	{"nouvelle", "nouvel"},
	{"seulement", "seulmen"},
	{"vrai", "vré"},
	{"seulement", "seulmen"},
	{"prendre", "prendr"},
	{"comme", "km"},
	{"moment", "momen"},
	{"ça fait", "sfé"},
	{"surtout", "surtt"},
	{"gueule", "gueul"},
	{"partout", "partt"},
}

func Smsify(input string) string {
	input = strings.ToLower(input)
	for _, pair := range mapping {
		input = strings.Replace(input, pair[0], pair[1], -1)
	}
	return input
}

// FIXME: implement "Unsmsify" :)
