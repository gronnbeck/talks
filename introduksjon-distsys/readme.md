# Introduksjon til distribuerte systemer 1

## Fallacies of distributed computing

  1. The network is reliable.
  1. Latency is zero.
  1. Bandwidth is infinite.
  1. The network is secure.
  1. Topology doesn't change.
  1. There is one administrator.
  1. Transport cost is zero.
  1. The network is homogeneous.


## The network is reliable

Anta at du har 2 systemer la oss kalle dem A og B. De er igjen avhengig av
andre systemer og andre systemer er avhengig av dem. Hva slags feil kan oppstå
hvis disse to får en nettverkspartisjon?

```
To noder A og B

    A         B
    o ------- o


    A         B
    o ---x--- o

```

  * Hva skjer hvis A er en applikasjon og B er en database eller en tjeneste?
    * Skal hele applikasjonen falle ned?
  * Hva skjer hvis A er en, ikke persisterende, bus/adapter for en applikasjon B?
    * Går data da tapt?


### Network partition vs node failure

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

Det er ingen måten for node ``B`` å vite at ``A`` er oppe eller om det kun
er nettverket mellom dem som er nede.

Hva så? Hva er problemet? Om du skal designe et distribuert system så må du
sørge for at en antatt død node ikke gjør noe den ikke skal gjøre.
Hvis du ikke sørger for dette ender du opp med å skape ditt eget "bysantisk
general problem" ved at din "døde" node kan skape problemer for f.eks.
data-strukturen din.

Takeaway:

## Time and ordering assumptions

En node i et distribuert system kan ikke vite den globale (sanne) rekkefølgen
på hendelser.

```
La A og B være 2 noder/kompoenenter av et distribuert system.

La a, b og c være hendelser

A ----b--------

B ----a-----c--
```

Vi kan ikke bruke "tiden" som vi bruker den. Siden vi ikke kan garantere at
klokka på hver maskin er synkronisert.

Ved hjelp av logiske klokker kan man vite om en hendelse skjer før en annen
hvis og bare hvis en hendelse fører til en annen. F.eks. hvis b fører til c
vet vi at b må ha vært før c. Men hvilken hendelse var først av a og b?

```

A ----b--------
         \
          \
           \
B ----a-----c--

```

## Distrbuert arkitektoniske mønstre

```
3-tier architecture:

   Presentation
       |
     Logic
       |
    Storage

    Frontend   <----\
       |              App
   Middleware  <----/
       |
    Database
```

### Tjenester som snakker med hverandre
Aka SOA.

Microservices: Små tjenester som snakker med hverandre aka SOA 2.0

  * Synkron distribuert arkitektur
    * n-tier arkitektur
  * Asynkron distrbuerte arkitektur
    * Smart bus, dumb applications
    * Dumb bus, smart applications
    * CQRS
  * Distribuerte Konsepter
    * Arbeidere
    * Datastrukturer
      * Køer
      * Hasher
      * Set
      * [etc][distributed-data-structures]
    * CAP ([stjålet herfra][distsys-4-fun-and-profit]):
     * CA (consistency + availability). Examples include full strict quorum protocols, such as two-phase commit.
     * CP (consistency + partition tolerance). Examples include majority quorum protocols in which minority partitions are unavailable such as Paxos.
     * AP (availability + partition tolerance). Examples include protocols using conflict resolution, such as Dynamo.



## Why now?

DevOps bevegelsen! Dev tar mer ansvar for sine egne applikasjoner. Ops sørger for at dev får lov til og kan klare dette  

Flere har bevist at det er mulig å få til en Tjenestearkitektur, men det krever endringer i organisasjon.

Vi har lært at SOA ikke kommer gratis. I disse dager kaller Martin Fowler det for [Microservice Premium](http://martinfowler.com/bliki/MicroservicePremium.html)


## SOA v1.0 died


[SOA is dead](http://apsblog.burtongroup.com/2009/01/soa-is-dead-long-live-services.html):
  * The acronym got in the way. People forgot what SOA stands for. They were too wrapped up in silly technology debates (e.g., “what’s the best ESB?” or “WS-* vs. REST”), and they missed the important stuff: architecture and services.
  * The small select group of organizations that has seen spectacular gains from SOA did so by treating it as an agent of transformation. In each of these success stories, SOA was just one aspect of the transformation effort. And here’s the secret to success: SOA needs to be part of something bigger.

[did-soa-die](http://www.kavistechnology.com/blog/did-soa-die-or-do-we-just-suck-at-architecture/):

  * Finally, the hype was just too much for those charged with SOA to resist. Projects selected the technology first, then the approach and architecture. That’s completely backwards.  

### Ikke lest, men burde kanskje

  * [Top 10 reasons why people are making SOA fail]( http://www.cio.com/article/2434865/service-oriented-architecture/top-10-reasons-why-people-are-making-soa-fail.html)


## Hva finnes?

Hva finnes av distribuerte systemer:

  * Prosessering
  * Lagring
  * etc

[distributed-data-structures]: http://www.gridgain.com/developer-central/in-memory-data-fabric/in-memory-data-grid/distributed-data-structures/
[distsys-4-fun-and-profit]: http://book.mixu.net/distsys/single-page.html
