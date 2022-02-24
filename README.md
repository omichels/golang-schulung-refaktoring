## Refaktoring


### Code Smell


Die Funktion CalculateTotalSum greift direkt auf die Attribute von
Employee zu und berechnet "selbst".

Was passiert, wenn der Mitarbeiter einen Bonus bekommt?

 sum = sum + ((v.hourlyPay * v.clockedHours) + v.Bonus)

 super hässlich, was ist wenn das Attribut nicht korrekt initialisiert ist mit 0