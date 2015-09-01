
# Eksempler

Noen eksempler på kjente distribuerte systemer
* Email, Slack, etc
* MMO-spill
* SOA:
  * SOA v1.0
  * SOA v2.0 (Microservices)

---

# Problems

Feil:
* Nettverks
* Node
* Race-conditions
* Synkronisering / konsistens

---

# Time and ordering

En node i et distribuert system kan ikke vite den globale (sanne) rekkefølgen
på hendelser.

La A og B være 2 noder/kompoenenter av et distribuert system.

La a, b og c være hendelser

```

A ----b--------

B ----a-----c--
```

```

A ----b--------
         \
          \
           \
B ----a-----c--

```

???

Vi kan ikke bruke "tiden" som vi bruker den. Siden vi ikke kan garantere at
klokka på hver maskin er synkronisert.

Ved hjelp av logiske klokker kan man vite om en hendelse skjer før en annen
hvis og bare hvis en hendelse fører til en annen. F.eks. hvis b fører til c
vet vi at b må ha vært før c. Men hvilken hendelse var først av a og b?

---

# The network is reliable

```
To noder A og B

    A         B
    o ------- o


    A         B
    o ---x--- o

```

???

Anta at du har 2 systemer la oss kalle dem A og B. De er igjen avhengig av
andre systemer og andre systemer er avhengig av dem. Hva slags feil kan oppstå
hvis disse to får en nettverkspartisjon?

  * Hva skjer hvis A er en applikasjon og B er en database eller en tjeneste?
    * Skal hele applikasjonen falle ned?
  * Hva skjer hvis A er en, ikke persisterende, bus/adapter for en applikasjon B?
    * Går data da tapt?


---

# Nettverkspartisjon vs node feil
```
Two nodes talks through a communication link:
    A         B
    o ------- o

Node failure:

    A         B
    x ------- o

Network partition:

    A         B
    o ---x--- o
```

???

Det er ingen måten for node ``B`` å vite at ``A`` er oppe eller om det kun
er nettverket mellom dem som er nede.

Hva så? Hva er problemet? Om du skal designe et distribuert system så må du
sørge for at en antatt død node ikke gjør noe den ikke skal gjøre.
Hvis du ikke sørger for dette ender du opp med å skape ditt eget "bysantisk
general problem" ved at din "døde" node kan skape problemer for f.eks.
data-strukturen din.

---

# CAP

---

# Konsensus

---

# Fallacies of distributed programming

  1. The network is reliable.
  1. Latency is zero.
  1. Bandwidth is infinite.
  1. The network is secure.
  1. Topology doesn't change.
  1. There is one administrator.
  1. Transport cost is zero.
  1. The network is homogeneous.

---

class: center, middle

# Tekniske endringer over tid

---

class: center, middle

#  Remote OO

---

class: center, middle

# Java RMI

![Java RMI](img/rmi.png)

???

* Sender objekter mellom maskiner
* Objekter deles via en "broker" altså en server
* Seralizerer alt i et objekt. Tom methodene.
* Hva skjer hvis man endrer på et objekt.

---

# Tekniske endringer

* SMTP, enkle meldinger
* Remote Object Orientet programming
  * Java RMI
  * Corba
* SOAP
* WebRPC
* JSON API ("REST")
* REST
* HATEOAS

---
