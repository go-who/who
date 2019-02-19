package main

import (
	"fmt" 
	"gopkg.in/who.v1"
	)

/*func WeAre(team string) {
//fmt.Printf("\n\n%#v\n", team)
fmt.Printf("\n\nTakim Uyeleri: %s", team)
}*/

func main(){
team := "\n Ad: Cemil Kiyak,\n Meslek: Mekatronik Muhendisi,\n Dogum yili:1987,\n E-posta: cemilbakikiyak@gmail.com,\n Firma: Devr-i Robotik Arge Ltd. Sti.,\n WEB: www.devrirobotik.com"

	
//fmt.Printf("%#v\n", team)
fmt.Printf("Takim uyesi bilgileri: %s", team)

	
//WeAre(team)
who.WeAre(team)

	
fmt.Println(`
                                         :tI
                                         t@:
                                        .@@
      .1@@@@t.    :f@@@@:     :t@@@@t:  :@1    1@@.
     f@@:  .@@: I@@:   f@t  I@@:   .f@t t@: .f@@.
    1@t     t@I:@@    .f@1 f@:      I@t.@f:@@f.
    @@:    .@@.1@@@@@@@f..@@.:@@@@@@@@::@@@f1@@@.
   :@@     I@t @@.      I@t f@1    :@@ t@:    I@t
   t@@1. .t@@. @@t.  .1@@:  @@1  .1@@..@t     t@I
  .@@:t@@@t.    :f@@@@1.     :@@@@1.  :@:    .@@.
  :@f
  t@:
  11.
   https://peak.com/jobs   
`)
}
