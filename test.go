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
var pozycja_wezaX=4
var pozycja_wezaY=10
var clear map[string]func()
var historiaWspolzednejX[10000] int
var historiaWspolzednejY[10000] int
var dlugosc_weza=3
var ile=dlugosc_weza
var pozycja_smakolykaX=13
var pozycja_smakolykaY=10
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
            plansza[pozycja_wezaY][pozycja_wezaX]='■'
          fmt.Printf("%c", plansza[i][j])
    }
    fmt.Println(" ")
}

c <- true
}
func sterujWezem() {

if kierunek=='D'{//down
  pozycja_wezaY++
}
if kierunek=='U'{//up
  pozycja_wezaY--
}
if kierunek=='L'{//left
  pozycja_wezaX--
}
if kierunek=='R'{//right
  pozycja_wezaX++
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
  if pozycja_wezaX==pozycja_smakolykaX && pozycja_wezaY==pozycja_smakolykaY{
    dlugosc_weza++
    wynik++
    pozycja_smakolykaX=randomInt(1,65)
    pozycja_smakolykaY=randomInt(1,25)
    plansza[pozycja_smakolykaY][pozycja_smakolykaX]='$'
  }

}

func main() {
rand.Seed(time.Now().UnixNano())
c := make(chan bool)
plansza=stworzPlansze()


for{
      ile++
      historiaWspolzednejX[ile]=pozycja_wezaX
      historiaWspolzednejY[ile]=pozycja_wezaY

sprawdzCzyZjadl()



go sterujWezem()

plansza[historiaWspolzednejY[ile-dlugosc_weza]][historiaWspolzednejX[ile-dlugosc_weza]]=' '

go rysujPlansze(c)
time.Sleep((1/2)*time.Second)
<- c

      CallClear()
}
}
