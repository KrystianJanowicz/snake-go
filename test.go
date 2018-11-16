package main

import (
  "fmt"
  "os"
  "os/exec"
  "math/rand"
  "runtime"
  "time"
  "github.com/eiannone/keyboard"
)
/* TO DO:

ranking
poprawa czytelnosci kodu

*/
var plansza [25][65]rune
var pozycjaWezaX=4
var pozycjaWezaY=10
var clear map[string]func()
var historiaWspolzednejX[10000] int
var historiaWspolzednejY[10000] int
var dlugoscWeza=3
var ile=dlugoscWeza
var pozycjaSmakolykaX=13
var pozycjaSmakolykaY=10
var kierunek rune='R'
var wynik int

func init() {
    clear = make(map[string]func())
    clear["linux"] = func() {
        cmd := exec.Command("clear") //linux
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //windows
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}
func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS = linux, windows
    if ok {
        value()
    } else {
        panic("nieznana platforma")
    }
}
func stworzPlansze() [25][65]rune {

var plansza [25][65]rune

for j:=0; j<65; j++{
  for i:=0; i<25; i++{
    plansza[i][j]=' '
  }
}

for i:=0; i<25; i++{
  plansza[i][0]='!'
}
for i:=0; i<25; i++{
  plansza[i][64]='!'
}
for i:=0; i<65; i++{
  plansza[0][i]='!'
}
for i:=0; i<65; i++{
  plansza[24][i]='!'

}
return plansza
}
func rysujPlansze(c chan bool) {
  fmt.Printf("Wynik: %d\n", wynik)
  for i:=0; i<25; i++{
    for j:=0; j<65; j++{
            plansza[pozycjaWezaY][pozycjaWezaX]='■'
          fmt.Printf("%c", plansza[i][j])
    }
    fmt.Println(" ")
}

c <- true
}
func sterujWezem() {

if kierunek=='D'{//down
  pozycjaWezaY++
}
if kierunek=='U'{//up
  pozycjaWezaY--
}
if kierunek=='L'{//left
  pozycjaWezaX--
}
if kierunek=='R'{//right
  pozycjaWezaX++
}

char, _, err:=keyboard.GetSingleKey()
if (err != nil) {
    panic(err)
}
    if char=='s'{
      kierunek='D'
    }
    if char=='w'{
      kierunek='U'
    }
    if char=='a'{
      kierunek='L'
    }
    if char=='d'{
      kierunek='R'
    }
}
func sprawdzCzyZjadl() {
  if pozycjaWezaX==pozycjaSmakolykaX && pozycjaWezaY==pozycjaSmakolykaY{
    dlugoscWeza++
    wynik++
    pozycjaSmakolykaX=randomInt(1,65)
    pozycjaSmakolykaY=randomInt(1,25)
    plansza[pozycjaSmakolykaY][pozycjaSmakolykaX]='$'
  }

}
func sprawdzCzyUmarl() bool{

  for i:=0; i<25; i++{
  if pozycjaWezaY==i && pozycjaWezaX==0 {
    return true
  }
}
  for i:=0; i<25; i++{
  if pozycjaWezaY==i && pozycjaWezaX==64{
      return true
  }
}
  for i:=0; i<65; i++{
    if pozycjaWezaY==0 && pozycjaWezaX==i{
      return true
  }
}
  for i:=0; i<65; i++{
    if pozycjaWezaY==24 && pozycjaWezaX==i{
      return true
    }
  }
  return false
}
func main() {
rand.Seed(time.Now().UnixNano())
c := make(chan bool)
plansza=stworzPlansze()


for{
      ile++
      historiaWspolzednejX[ile]=pozycjaWezaX
      historiaWspolzednejY[ile]=pozycjaWezaY

sprawdzCzyZjadl()
if sprawdzCzyUmarl() == true {
break
}



go sterujWezem()

plansza[historiaWspolzednejY[ile-dlugoscWeza]][historiaWspolzednejX[ile-dlugoscWeza]]=' '

go rysujPlansze(c)
time.Sleep((1/2)*time.Second)
<- c

      CallClear()
}
fmt.Printf("przykro mi ale sie wywaliles, twój wynik: %d", wynik)

}
