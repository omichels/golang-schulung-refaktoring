## Refaktoring


### Code Smell


Die Funktion CalculateTotalSum greift direkt auf die Attribute von
Employee zu und berechnet "selbst".

Was passiert, wenn der Mitarbeiter einen Bonus bekommt, 
dafür dass er mehr als 38h im letzten Monat gearbeitet hat ?

 if (v.clockedHours > 38) {
   sum = sum + ((v.hourlyPay * v.clockedHours) + v.Bonus)
 } else {
   sum = sum + (v.hourlyPay * v.clockedHours)
 }

 super hässlich, was ist, wenn das Attribut nicht korrekt initialisiert ist mit 0,
 was, wenn in anderen Niederlassungen 40h schon Standard ist, 
 was, wenn die 38h als Regel geändert werden, 
 was, wenn ...
 

Wie sieht's denn besser aus?

Hollywood-Principle: Rufen Sie uns nicht an, wir rufen Sie an!

Wir sagen der Entität also, dass es eine Berechnung durchführen soll



Stufe 2:
Es gibt die Typen Employee und ChinaEmployee
ChinaEmployee erhält den Bonus erst, wenn mehr als 100 Arbeitsstunden gebucht wurden.

Die Variablität ist aber nun gekapselt, Änderungen an den Arbeitsverträgen schlagen nicht
mehr bis "ganz oben" in die Berechnung von CalculateTotalSum durch.
Hier gibt's niemals
 if 
  else
   if 
    else
     if 
      else


## Code Smell DRY 
Do not repeat yourself

Type Embedding to the rescue

before                                  after

type ChinaEmployee struct {             type ChinaEmployee struct {
  Id           int                        Emp Employee
  HourlyPay    int                      }
  ClockedHours int
  Bonus        int
}






