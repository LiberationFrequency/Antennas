
![Preview](https://raw.githubusercontent.com/LiberationFrequency/Antennas/master/4nec2/modeling/70cm-rhombus_50_vert_awg22/70cm-rhombus_50_awg22_crop.png)


Wer hat's gemacht?  
-----------------------------------
LiberationFrequency  
https://github.com/LiberationFrequency/Antennas/
  
Inspiriert durch DK7ZB / Funkamateur Magazin Februar 2013 S. 182-185  
  
Was ist das?  
-----------------------------------
Das ist eine Ganzwellenschleife als Rhombus, ein gestocktes/gestauchtes/gestrecktes(?) Quadrat, was es im rechten Winkel ist.  
Der Fußpunktwiderstand liegt bei 50 Ω, und bleibt in fast jeder Situation stabil.  
Was sich ändert ist die Gewinnerwartung, da kann man nicht so viel drauf geben. Das wird stark von der Höhe in Wellenlängen über dem Grund beeinflusst.  


Wie wird's gemacht?  
-----------------------------------

Vorbereitung:  
Ein Stück AWG22 nehmen und abhängen, damit sich die Länge später weniger beim strammziehen setzt.  
  
Beim reinen Empfang muss man nicht immer so genau sein, beim Senden aber schon.  
Um Störungen durch entstehende Mantelwellen zu vermeiden, sollte man eine 1:1 Übesetzung benutzen, oder zumindest eine Luftspule mit 10 Windungen in das Koaxialkabel machen. 

Das AWG22-Kabel ist sehr dünn, das wird eine instabile, ultra-low-cost Antenne.
Zuerst schneiden wir die Diagonalen zurecht. Ein bisschen was länger lassen, um im übernächsten Schritt Einkerbungen zur besseren Führung des Kabeln feilen zu können.  
  
Die Mitte der beiden Diagonalen bestimmen und zusammen befestigen, nageln, schrauben, knoten, wie auch immer.  

Enden zur besseren Führung des Kabels mit einer Feile einkerben.  
  
Berechnen:
Ich habe einen kleinen Rechner in Golang geschrieben.  
  
./antcalc --alpha=53.0 --factor=1.094842 --freq=435.0  
Frequency in MHz:              435.00000  
Wavelenght in m:               0.68918  
Speed of light in Mm/s:        299.792458  
Extension/shortening factor:   1.09484  
Perimeter in m incl. factor:   0.75454  
Sidelenght in m:               0.18864  
Alpha angle in degree:         53.00000  
Beta angle in degree:          127.00000  
Horizontal diagonal in m:      0.33763  
Vertical diagonal in m:        0.16834  
Surface area in m²:            0.02842  
Apothem in m:                  0.07533  
  

Who did it?  
-----------------------------------
LiberationFrequency  
https://github.com/LiberationFrequency/Antennas/
  
Inspired by DK7ZB / Funkamateur magazine February 2013 p. 182-185 in german.  
  
What's that?
-----------------------------------
This is a full wave loop as a rhombus, a stacked/compressed/stretched(?) square, which it is at right angle.
The foot point resistance is 50 Ω, and remains stable in almost any situation.
What changes is the profit expectation, there is not so much you can put on it. This is strongly influenced by the height in wavelengths above the ground.


How to?
-----------------------------------
Preparation:
Take a piece of AWG22 and hang it up so that the length will settle less when tightening later.

You don't always have to be so precise when you're just receiving, but you have to do when you're sending.
In order to avoid disturbances caused by sheath current waves, a 1:1 ratio should be used, or at least an air coil with 10 turns in the coaxial cable.

The AWG22 cable is very thin, which becomes an unstable, ultra-low-cost antenna.
First we cut the diagonals to size. Leave a little longer to be able to file notches in the next but one step to better guide the cable.

Determine the center of the two diagonals and fasten together, nail, screw, knot, however.

Notch the ends with a file to guide the cable better.  
  
Calculate:
I wrote a little calculator in Golang.  
  
./antcalc --alpha=53.0 --factor=1.094842 --freq=435.0  
Frequency in MHz:              435.00000  
Wavelenght in m:               0.68918  
Speed of light in Mm/s:        299.792458  
Extension/shortening factor:   1.09484  
Perimeter in m incl. factor:   0.75454  
Sidelenght in m:               0.18864  
Alpha angle in degree:         53.00000  
Beta angle in degree:          127.00000  
Horizontal diagonal in m:      0.33763  
Vertical diagonal in m:        0.16834  
Surface area in m²:            0.02842  
Apothem in m:                  0.07533  
  