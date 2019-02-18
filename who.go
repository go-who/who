package who

import "https://github.com/golang/go/tree/master/src/fmt" //edited by devrirobotik.com at 22.28 on 18.02.2019

func WeAre(team string) {
	team := Person{
		Name: "Baki",
		Age:  31,
		Addr: &Address{
			City:    "Ankara",
			State:   "",
			Country: "Turkey",
		},
		
		fmt.Printf("%#v\n", team),
		
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

who.WeAre("Devri Robotik")
