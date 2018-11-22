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
/*
commit zawiera:
-wąż ginie od zjedzenia samego siebie-
-smakołyk przyspieszający bieg węża-
*/
const szerokoscPlanszy=65
const dlugoscPlanszy=25
var plansza [dlugoscPlanszy][szerokoscPlanszy]rune
var pozycjaWezaX=10
var pozycjaWezaY=10
var clear map[string]func()
var historiaWspolzednejX[10000] int
var historiaWspolzednejY[10000] int
var dlugoscWeza=3
var ile=dlugoscWeza
var pozycjaSmakolykaX=10
var pozycjaSmakolykaY=10
var kierunek rune='R'
var wynik=-1
var wybor int
var pozycjaBoosteraX int
var pozycjaBoosteraY int
var czyBoosterIstnieje=false
var czasBoostera=0

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
func reset() {
wynik=-1
pozycjaSmakolykaX=10
pozycjaSmakolykaY=10
pozycjaWezaX=10
pozycjaWezaY=10
dlugoscWeza=3
ile=dlugoscWeza
kierunek='R'
}
func stworzPlanszeZRamka() [dlugoscPlanszy][szerokoscPlanszy]rune {

var plansza [dlugoscPlanszy][szerokoscPlanszy]rune

for j:=0; j<szerokoscPlanszy; j++{
  for i:=0; i<dlugoscPlanszy; i++{
    plansza[i][j]=' '
  }
}

for i:=0; i<dlugoscPlanszy; i++{
  plansza[i][0]='a'
}
for i:=0; i<dlugoscPlanszy; i++{
  plansza[i][szerokoscPlanszy-1]='s'
}
for i:=0; i<szerokoscPlanszy; i++{
  plansza[0][i]='d'
}
for i:=0; i<szerokoscPlanszy; i++{
  plansza[dlugoscPlanszy-1][i]='b'

}
return plansza
}
func stworzPlanszeBez() [dlugoscPlanszy][szerokoscPlanszy]rune {

var plansza [dlugoscPlanszy][szerokoscPlanszy]rune

for j:=0; j<szerokoscPlanszy; j++{
  for i:=0; i<dlugoscPlanszy; i++{
    plansza[i][j]=' '
  }
}

return plansza
}
func rysujPlansze(c chan bool) {
  fmt.Printf("Wynik: %d\n", wynik)
  for i:=0; i<dlugoscPlanszy; i++{
    for j:=0; j<szerokoscPlanszy; j++{
            plansza[pozycjaWezaY][pozycjaWezaX]='■'
          fmt.Printf("%c", plansza[i][j])
    }
    fmt.Println(" ")
}

c <- true
}
func sterujWezem() {
var nowyKierunek rune

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
      nowyKierunek='D'
    }
    if char=='w'{
      nowyKierunek='U'
    }
    if char=='a'{
      nowyKierunek='L'
    }
    if char=='d'{
      nowyKierunek='R'
    }
    if   kierunek=='U' && nowyKierunek=='D' || kierunek=='D' && nowyKierunek=='U' || kierunek=='L' && nowyKierunek=='R' || kierunek=='R' && nowyKierunek=='L'{
      nowyKierunek=kierunek
    } else  {
    kierunek=nowyKierunek
  }
}
func sprawdzCzyZjadl() {
  if pozycjaWezaX==pozycjaSmakolykaX && pozycjaWezaY==pozycjaSmakolykaY{
    dlugoscWeza++
    wynik++
    pozycjaSmakolykaX=randomInt(1,szerokoscPlanszy-1)
    pozycjaSmakolykaY=randomInt(1,dlugoscPlanszy-1)
    plansza[pozycjaSmakolykaY][pozycjaSmakolykaX]='$'
  }

}
func wygernerujISprawdzCzyZjadlBooster() {
  if ile>50 && czyBoosterIstnieje==false {
    pozycjaBoosteraX=randomInt(1,szerokoscPlanszy-1)
    pozycjaBoosteraY=randomInt(1,dlugoscPlanszy-1)
    plansza[pozycjaBoosteraY][pozycjaBoosteraX]='B'
    czyBoosterIstnieje=true
  }
    if pozycjaWezaX==pozycjaBoosteraX && pozycjaWezaY==pozycjaBoosteraY{
      czyBoosterIstnieje=false
      czasBoostera=50
    }


}
func sprawdzCzyUmarlRamka() bool{

  for i:=0; i<dlugoscPlanszy; i++{
  if pozycjaWezaY==i && pozycjaWezaX==0 {
    return true
  }
}
  for i:=0; i<dlugoscPlanszy; i++{
  if pozycjaWezaY==i && pozycjaWezaX==szerokoscPlanszy-1{
      return true
  }
}
  for i:=0; i<szerokoscPlanszy; i++{
    if pozycjaWezaY==0 && pozycjaWezaX==i{
      return true
  }
}
  for i:=0; i<szerokoscPlanszy; i++{
    if pozycjaWezaY==dlugoscPlanszy-1 && pozycjaWezaX==i{
      return true
    }
  }
  //waz zjada sam siebie
  for i:=1;i<=dlugoscWeza;i++ {
      if pozycjaWezaX==historiaWspolzednejX[ile-i] && pozycjaWezaY==historiaWspolzednejY[ile-i] {
      return true
      }
    }
  return false
}
func sprawdzCzyUmarlBez() bool{
//przechodzenie weza przez sciany
if pozycjaWezaX==szerokoscPlanszy-1{
  pozycjaWezaX=1
}
if pozycjaWezaX==0{
  pozycjaWezaX=szerokoscPlanszy-1
}
if pozycjaWezaY==dlugoscPlanszy-1{
  pozycjaWezaY=1
}
if pozycjaWezaY==0{
  pozycjaWezaY=dlugoscPlanszy-1
}
//waz zjada sam siebie
for i:=1;i<=dlugoscWeza;i++ {
    if pozycjaWezaX==historiaWspolzednejX[ile-i] && pozycjaWezaY==historiaWspolzednejY[ile-i] {
    return true
    }
  }
return false
}

func start() {

fmt.Println("SNAKE. wybierz 1 aby grać bez ramki, 0 aby grać z")
fmt.Scan(&wybor)

rand.Seed(time.Now().UnixNano())
c := make(chan bool)

if wybor==1{
plansza=stworzPlanszeBez()
}

if wybor==0{
plansza=stworzPlanszeZRamka()
}

for{
      ile++
      historiaWspolzednejX[ile]=pozycjaWezaX
      historiaWspolzednejY[ile]=pozycjaWezaY

sprawdzCzyZjadl()

if wybor==1{
  if sprawdzCzyUmarlBez() == true {
  break
  }
}
if wybor==0{
  if sprawdzCzyUmarlRamka() == true {
  break
  }
}
go sterujWezem()
plansza[historiaWspolzednejY[ile-dlugoscWeza]][historiaWspolzednejX[ile-dlugoscWeza]]=' '
go rysujPlansze(c)

wygernerujISprawdzCzyZjadlBooster()

if czasBoostera>0{
  time.Sleep(40000000*time.Nanosecond)
  czasBoostera--
}
if czasBoostera<=0{
  time.Sleep(80000000*time.Nanosecond)
}

<- c
CallClear()
}
}

func main() {

for{
start()
fmt.Printf("przykro mi ale sie wywaliles, twój wynik: %d\n", wynik)
reset()
}
}
