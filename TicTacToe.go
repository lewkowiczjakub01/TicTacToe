package main

import (
  "fmt"
  "os"
  "time"
  "math/rand"
)


func main() {
  Pola := [9]string{"A1 ","A2 ","A3 ","B1 ","B2 ","B3 ","C1 ","C2 ","C3 "} //Pola planszy
  var x int = 0 // do obslugi menu
  var enter string //tylko po to by wychodzic z "Zasad gry"
  for i:=0 ; i<100 ; i++ { // aby powracalo do menu 
  fmt.Printf("\033[H\033[2J \033[2;15f")
  fmt.Println(" \033[32m \x1B[4m KÓŁKO I KRZYŻYK \x1B[24m \033[97m  ")
  fmt.Println("\033[3;15f1. Nowa gra z komputerem")
  fmt.Println("\033[4;15f2. Nowa gra z innym graczem")
  fmt.Println("\033[5;15f3. Zasady gry ")
  fmt.Println("\033[6;15f0. Wyjście ")
  fmt.Scanf("%d", &x)

  if x==1 {
    Rozpocznij_grę_z_komputerem(Pola)
  } else if x==2 {
    Rozpocznij_grę_z_innym_graczem(Pola)
  }  else if x==3 {
    fmt.Printf("\033[H\033[2J ")
    fmt.Println("\033[2;10f 1. Gra toczy się na kwadratowej planszy 3x3")
    fmt.Println("\033[3;10f 2. Naprzemiennie stawia się symbole X i O")
    fmt.Println("\033[4;10f 3. Wygrywa gracz który posiada 3 swoje symbole\n  \033[5;10f     w linii, rzędzie, na ukos lub wspak")
    fmt.Println("\n\nNaciśnij enter by wrócic do menu")
    fmt.Scanf("%s",&enter)
    fmt.Printf("\033[H\033[2J")
  } else if x==0 {
    os.Exit(0)
  }

}
}

func plansza (Pola [9]string) { //rysowanie planszy
  fmt.Printf("\033[H\033[2J \033[2;15f")
  fmt.Printf(" %s\033[36m║\033[97m%s\033[36m║\033[97m%s\n", Pola[0],Pola[1],Pola[2])
  fmt.Println(" \033[3;15f \033[36m═══╬═══╬═══\033[97m")
  fmt.Printf(" \033[4;15f %s\033[36m║\033[97m%s\033[36m║\033[97m%s\n",Pola[3],Pola[4],Pola[5])
  fmt.Println(" \033[5;15f \033[36m═══╬═══╬═══\033[97m")
  fmt.Printf("\033[6;15f %s\033[36m║\033[97m%s\033[36m║\033[97m%s\n\n\n",Pola[6],Pola[7],Pola[8])
}

func Rozpocznij_grę_z_komputerem (Pola[9] string) {
  var pole string //wartosc podana przez gracza
  var w int =2 // zmienia sie na w=1 jezeli ktos wygrywa
  var los int =0 //daje losowe pole komputerowi jezeli potrzeba
  for i := 0 ; i<10 ; i++ { // 9 ruchów na planszy + 10 sie uaktywnia jezeli jest remis
  plansza(Pola)

  if i%2 == 0 && i<9 { //ogranicza do 9 bo 9 pol na planszy
  fmt.Printf("\033[33mRuch gracza\n Wybierz pole: \033[97m")
  fmt.Scanf("%s", &pole)
  pole += " " //dodaje spacje gdyz w wartosciach w tabeli sa z spacjami
  if Pola[0] == pole && Pola[1] != " X " && Pola[1] != " O "{ //sprawdzam czy podane pole nie jest juz zajete
    Pola[0] = " X "
  } else if Pola[1] == pole && Pola[1] != " X" && Pola[1] != " O "{
    Pola[1] = " X "
  } else if Pola[2] == pole && Pola[2] != " X " && Pola[2] != " O "{
    Pola[2] = " X "
  } else if Pola[3] == pole && Pola[3] != " X " && Pola[3] != " O "{
    Pola[3] = " X "
  } else if Pola[4] == pole && Pola[4] != " X " && Pola[4] != " O "{
    Pola[4] = " X "
  } else if Pola[5] == pole && Pola[5] != " X " && Pola[5] != " O "{
    Pola[5] = " X "
  } else if Pola[6] == pole && Pola[6] != " X " && Pola[6] != " O "{
    Pola[6] = " X "
  } else if Pola[7] == pole && Pola[7] != " X " && Pola[7] != " O "{
    Pola[7] = " X "
  } else if Pola[8] == pole && Pola[8] != " X " && Pola[8] != " O "{
    Pola[8] = " X "
  } else { //daje 2 szanse na podanie prawidlowej wartosci
    fmt.Println("\033[31mPodano złą wartość. \n Proszę podać inną: \033[97m")
    var pole2 string
    fmt.Scanf("%s", &pole2)
    pole2 += " "
        if Pola[0] == pole2 && Pola[1] != " X " && Pola[1] != " O "{
           Pola[0] = " X "
        } else if Pola[1] == pole2 && Pola[1] != " X " && Pola[1] != " O "{
           Pola[1] = " X "
        } else if Pola[2] == pole2 && Pola[2] != " X " && Pola[2] != " O "{
           Pola[2] = " X "
        } else if Pola[3] == pole2 && Pola[3] != " X " && Pola[3] != " O "{
           Pola[3] = " X "
        } else if Pola[4] == pole2 && Pola[4] != " X " && Pola[4] != " O "{
           Pola[4] = " X "
        } else if Pola[5] == pole2 && Pola[5] != " X " && Pola[5] != " O "{
           Pola[5] = " X "
        } else if Pola[6] == pole2 && Pola[6] != " X " && Pola[6] != " O "{
           Pola[6] = " X "
        } else if Pola[7] == pole2 && Pola[7] != " X " && Pola[7] != " O "{
           Pola[7] = " X "
        } else if Pola[8] == pole2 && Pola[8] != " X " && Pola[8] != " O "{
           Pola[8] = " X "
        } else { //jezeli 2 razy poda sie zla wartosc to program sie zamyka z errorem
          fmt.Println("\033[31mError: Podawanie nieodpowiednich wartości\033[97m")
          os.Exit(0)
        }
  }
  plansza(Pola) //zawsze po ruchu rysuje plansze na nowo i sprawdzam czy ktos wygrywa
  w = CzyWygral(Pola) //jezeli ktos wygrywa to zwraca wartosc 1
  if w == 1 {
     fmt.Println("\033[32mWygrał gracz pierwszy !")
     fmt.Println("Powrót do menu za 10 sekund\033[97m")
     time.Sleep(10 * time.Second)
     break
  }
  }  else if i%2 != 0 && i<9 { //AI komputera
  if i == 1 && Pola[4] == "B2 " { //jezeli gracz na pocatek nie wybral srodkowego pola to komputer je zajmie
    Pola[4] = " O "
  } else if Pola[0] == " X " && Pola[1] == " X " && Pola[2] != " O "{ //komputer pierw sprawdza czy gracz nie jest o jedne krok od zwyciestwa i blokuje go jezeli jest
    Pola[2] = " O "
  } else if Pola[1] == " X " && Pola[2] == " X " && Pola[0] != " O "{
    Pola[0] = " O "
  } else if Pola[3] == " X " && Pola[4] == " X " && Pola[5] != " O "{
    Pola[5] = " O "
  } else if Pola[4] == " X " && Pola[5] == " X " && Pola[3] != " O "{
    Pola[3] = " O "
  } else if Pola[0] == " X " && Pola[2] == " X " && Pola[1] != " O "{
    Pola[1] = " O "
  } else if Pola[6] == " X " && Pola[7] == " X " && Pola[8] != " O "{
    Pola[8] = " O "
  } else if Pola[7] == " X " && Pola[8] == " X " && Pola[6] != " O "{
    Pola[6] = " O "
  } else if Pola[6] == " X " && Pola[8] == " X " && Pola[7] != " O "{
    Pola[7] = " O "
  } else if Pola[0] == " X " && Pola[3] == " X " && Pola[6] != " O "{
    Pola[6] = " O "
  } else if Pola[3] == " X " && Pola[6] == " X " && Pola[0] != " O "{
    Pola[0] = " O "
  } else if Pola[0] == " X " && Pola[6] == " X " && Pola[3] != " O "{
    Pola[3] = " O "
  } else if Pola[1] == " X " && Pola[4] == " X " && Pola[7] != " O "{
    Pola[7] = " O " 
  } else if Pola[4] == " X " && Pola[7] == " X " && Pola[1] != " O "{
    Pola[1] = " O " 
  } else if Pola[2] == " X " && Pola[5] == " X " && Pola[8] != " O "{
    Pola[8] = " O " 
  } else if Pola[5] == " X " && Pola[8] == " X " && Pola[2] != " O "{
    Pola[2] = " O " 
  } else if Pola[2] == " X " && Pola[8] == " X " && Pola[5] != " O "{
    Pola[5] = " O " 
  } else if Pola[0] == " X " && Pola[4] == " X " && Pola[8] != " O "{
    Pola[8] = " O " 
  } else if Pola[4] == " X " && Pola[8] == " X " && Pola[0] != " O "{
    Pola[0] = " O " 
  } else if Pola[2] == " X " && Pola[4] == " X " && Pola[6] != " O "{
    Pola[6] = " O " 
  } else if Pola[4] == " X " && Pola[6] == " X " && Pola[2] != " O "{
    Pola[2] = " O " 
  } else { //jezeli nie jest to losuje pole na ktorym stawia kółko
    rand.Seed(time.Now().UnixNano())
    los = rand.Intn(8)
    if Pola[los] != " X " && Pola[los] != " O " {
      Pola[los] = " O "
    } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
      } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
    } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
    } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
    } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
    } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
    } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
    } else {
      rand.Seed(time.Now().UnixNano())
      los = rand.Intn(8)
       if Pola[los] != " X " && Pola[los] != " O " {
            Pola[los] = " O "
    } else{ //jezeli w tych kilku probach nie uda sie wybrac nie zajetego pola to bierze pierwsze dostepne
      if Pola[0] != " X " && Pola[0] != " O "{
        Pola[0] = " O "
      } else if Pola[1] != " X " && Pola[1] != " O " {
        Pola[1] = " O "
      } else if Pola[2] != " X " && Pola[2] != " O " {
        Pola[2] = " O "
      } else if Pola[3] != " X " && Pola[3] != " O " {
        Pola[3] = " O "
      } else if Pola[5] != " X " && Pola[5] != " O " {
        Pola[5] = " O "
      } else if Pola[6] != " X " && Pola[6] != " O " {
        Pola[6] = " O "
      } else if Pola[7] != " X " && Pola[7] != " O " {
        Pola[7] = " O "
      } else if Pola[8] != " X " && Pola[8] != " O " {
        Pola[8] = " O "
      }
    }
    }}}}}}}
    }
 
    
  }
  plansza(Pola)
  w = CzyWygral(Pola)
  if w == 1 {
     fmt.Println("\033[32mWygrał komputer !")
     fmt.Println("Powrót do menu za 10 sekund\033[97m")
     time.Sleep(10 * time.Second)
     break
  }
  } else { //po 9 ruchach jezeli nikt nie wygral to jest remis (dlatego for jest do 10 a nie do 9)
    fmt.Println("\033[33m Remis \nPowrót do menu za 10 sekund \033[97m ")
    time.Sleep(10 * time.Second)
  }
  }
}

func Rozpocznij_grę_z_innym_graczem (Pola[9] string) { //tu jest bardzo podobnie do tego co wyzej
  var pole string
  var w int =2
  for i := 0 ; i<10 ; i++ {
  plansza(Pola)

  if i%2 == 0 && i<9{
  fmt.Printf("\033[33mRuch gracza nr 1\n Wybierz pole: \033[97m")
  fmt.Scanf("%s", &pole)
  pole += " "
  if Pola[0] == pole && Pola[1] != " X " && Pola[1] != " O "{
    Pola[0] = " X "
  } else if Pola[1] == pole && Pola[1] != " X" && Pola[1] != " O "{
    Pola[1] = " X "
  } else if Pola[2] == pole && Pola[2] != " X " && Pola[2] != " O "{
    Pola[2] = " X "
  } else if Pola[3] == pole && Pola[3] != " X " && Pola[3] != " O "{
    Pola[3] = " X "
  } else if Pola[4] == pole && Pola[4] != " X " && Pola[4] != " O "{
    Pola[4] = " X "
  } else if Pola[5] == pole && Pola[5] != " X " && Pola[5] != " O "{
    Pola[5] = " X "
  } else if Pola[6] == pole && Pola[6] != " X " && Pola[6] != " O "{
    Pola[6] = " X "
  } else if Pola[7] == pole && Pola[7] != " X " && Pola[7] != " O "{
    Pola[7] = " X "
  } else if Pola[8] == pole && Pola[8] != " X " && Pola[8] != " O "{
    Pola[8] = " X "
  } else {
    fmt.Println("\033[31mPodano złą wartość. \n Proszę podać inną: \033[97m")
    var pole2 string
    fmt.Scanf("%s", &pole2)
    pole2 += " "
        if Pola[0] == pole2 && Pola[1] != " X " && Pola[1] != " O "{
           Pola[0] = " X "
        } else if Pola[1] == pole2 && Pola[1] != " X " && Pola[1] != " O "{
           Pola[1] = " X "
        } else if Pola[2] == pole2 && Pola[2] != " X " && Pola[2] != " O "{
           Pola[2] = " X "
        } else if Pola[3] == pole2 && Pola[3] != " X " && Pola[3] != " O "{
           Pola[3] = " X "
        } else if Pola[4] == pole2 && Pola[4] != " X " && Pola[4] != " O "{
           Pola[4] = " X "
        } else if Pola[5] == pole2 && Pola[5] != " X " && Pola[5] != " O "{
           Pola[5] = " X "
        } else if Pola[6] == pole2 && Pola[6] != " X " && Pola[6] != " O "{
           Pola[6] = " X "
        } else if Pola[7] == pole2 && Pola[7] != " X " && Pola[7] != " O "{
           Pola[7] = " X "
        } else if Pola[8] == pole2 && Pola[8] != " X " && Pola[8] != " O "{
           Pola[8] = " X "
        } else {
          fmt.Println("\033[31mError: Podawanie nieodpowiednich wartości\033[97m")
          os.Exit(0)
        }
  }
  plansza(Pola)
  w = CzyWygral(Pola)
  if w == 1 {
     fmt.Println("\033[32mWygrał gracz pierwszy !")
     fmt.Println("Powrót do menu za 10 sekund \nPowrót do menu za 10 sekund \033[97m")
     time.Sleep(10 * time.Second)
     break
  }
  } else if i%2!=0 && i<9 {
  fmt.Printf("\033[33mRuch gracza nr 2\n Wybierz pole: \033[97m")
  fmt.Scanf("%s", &pole)
  pole += " "
  if Pola[0] == pole && Pola[1] != " X " && Pola[1] != " O "{
    Pola[0] = " O "
  } else if Pola[1] == pole && Pola[1] != " X " && Pola[1] != " O "{
    Pola[1] = " O "
  } else if Pola[2] == pole && Pola[2] != " X " && Pola[2] != " O "{
    Pola[2] = " O "
  } else if Pola[3] == pole && Pola[3] != " X " && Pola[3] != " O "{
    Pola[3] = " O "
  } else if Pola[4] == pole && Pola[4] != " X " && Pola[4] != " O "{
    Pola[4] = " O "
  } else if Pola[5] == pole && Pola[5] != " X " && Pola[5] != " O "{
    Pola[5] = " O "
  } else if Pola[6] == pole && Pola[6] != " X " && Pola[6] != " O "{
    Pola[6] = " O "
  } else if Pola[7] == pole && Pola[7] != " X " && Pola[7] != " O "{
    Pola[7] = " O "
  } else if Pola[8] == pole && Pola[8] != " X " && Pola[8] != " O "{
    Pola[8] = " O "
  } else {
    fmt.Println("\033[31mPodano złą wartość. \n Proszę podać inną: \033[97m")
    var pole3 string
    fmt.Scanf("%s", &pole3)
    pole3 += " "
        if Pola[0] == pole3 && Pola[1] != " X " && Pola[1] != " O "{
           Pola[0] = " O "
        } else if Pola[1] == pole3 && Pola[1] != " X " && Pola[1] != " O "{
           Pola[1] = " O "
        } else if Pola[2] == pole3 && Pola[2] != " X " && Pola[2] != " O "{
           Pola[2] = " O "
        } else if Pola[3] == pole3 && Pola[3] != " X " && Pola[3] != " O "{
           Pola[3] = " O "
        } else if Pola[4] == pole3 && Pola[4] != " X " && Pola[4] != " O "{
           Pola[4] = " O "
        } else if Pola[5] == pole3 && Pola[5] != " X " && Pola[5] != " O "{
           Pola[5] = " O "
        } else if Pola[6] == pole3 && Pola[6] != " X " && Pola[6] != " O "{
           Pola[6] = " O "
        } else if Pola[7] == pole3 && Pola[7] != " X " && Pola[7] != " O "{
           Pola[7] = " O "
        } else if Pola[8] == pole3 && Pola[8] != " X " && Pola[8] != " O "{
           Pola[8] = " O "
        } else {
          fmt.Println("\033[31mError: Podawanie nieodpowiednich wartości\033[97m")
          os.Exit(0)
        }
  }
  plansza(Pola)
  w = CzyWygral(Pola)
  if w == 1 {
     fmt.Println("\033[32mWygrał gracz drugi !")
     fmt.Println("Powrót do menu za 10 sekund\033[97m")
     time.Sleep(10 * time.Second)
     break
  }
  } else {
    fmt.Println("\033[33m Remis \nPowrót do menu za 10 sekund \033[97m ")
    time.Sleep(10 * time.Second)
  }
  }

}

func CzyWygral (Pola[9] string) int { //ta funnkcja sprawdza czy ktos akurat nie wygrywa i jezeli wygrywa to zwraca wartosc 1
  if Pola[0] == Pola[1] && Pola[1] == Pola[2]{
    if Pola[0] == " X "{
      return 1
    } else {
        return 1
    }
  } else if Pola[3] == Pola[4] && Pola[4] == Pola[5]{
    if Pola[3] == " X "{
        return 1
    } else {
        return 1
    }
  } else if Pola[6] == Pola[7] && Pola[7] == Pola[8]{
    if Pola[6] == " X "{
        return 1
    } else {
        return 1
    }
  } else if Pola[0] == Pola[4] && Pola[4] == Pola[8]{
    if Pola[0] == " X "{
        return 1
    } else {
        return 1
    }
  } else if Pola[2] == Pola[4] && Pola[4] == Pola[6]{
    if Pola[2] == " X "{
        return 1
    } else {
        return 1
    }
  } else if Pola[0] == Pola[3] && Pola[3] == Pola[6]{
    if Pola[2] == " X "{
        return 1
    } else {
        return 1
    }
  } else if Pola[1] == Pola[4] && Pola[4] == Pola[7]{
    if Pola[2] == " X "{
        return 1
    } else {
        return 1
    }
  } else if Pola[2] == Pola[5] && Pola[5] == Pola[8]{
    if Pola[2] == " X "{
        return 1
    } else {
        return 1
    }
  }
  return 0
}
