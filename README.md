# MusicXmlGoParser
Converts MusicXml into a machine learning friendly format

### USAGE Example:
Generate a midi melody for a given key plus a 2 bar chord progression like so:<br> 
```java -jar chords-to-melody-generator.jar -chords=31^313*313*313*313*613*613*715*715*```

# Encoding format
Each bar consists of a 72 digit encoding
##### [key][chords][bar1][bar2][bar3][bar4]

##### KEY EXAMPLE
[F-sharp] = [62]

##### CHORDS EXAMPLE
[G-major-7 | eb minor] = [713501]

##### BAR EXAMPLE
Each beat in a bar: [NOTE-SHARP-OCATAVE-LIFECYCLE]<br>
[g-sharp-quarter | rest-quarter | d-quarter | e-eighth | f-eight] = [7140-7141-7141-7141--0000-0000-0000-0000--4140-4141-4141-4141--5140-5141-6140-6140]

##### NOTES: <br>
0 | Rest<br>
1 | a<br>
2 | b<br>
3 | c<br>
4 | d<br>
5 | e<br>
6 | f<br>
7 | g<br>

##### FLAT/NATURAL/SHARP
0 - FLAT |
1 - NATURAL |
2 - SHARP

##### OCTAVES
0-5

##### NOTE LIFE_CYCLE
0 - START |
1 - SUSTAIN

##### CHORD TYPES
0 | maj<br>
1 | min<br>
2 | dim<br>
3 | maj7<br>
4 | min7<br>
5 | dom7<br>
6 | min7b5<br>#### CHORD TYPE

##### START/SUSTAIN
0 - START | 
1 - SUSTAIN


 
