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
