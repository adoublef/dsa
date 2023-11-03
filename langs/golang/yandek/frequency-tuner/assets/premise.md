Since childhood, Maxim has been a pretty good musician and a jack-of-all-trades. He recently made a simple percussion musical instrument himself - a triangle. He needs to find out what is the frequency of the sound his instrument makes.

Maxim has a professional musical tuner, which can be used to play a note with a given frequency. Maxim acts in the following way: he turns on the tuner notes with different frequencies and for each note he uses his hearing to determine whether it is closer or farther to the sound produced by the triangle than the previous note. Since Maksim's hearing is absolute, he always determines it absolutely correctly.

Maxim showed you a record which has a sequence of frequencies you set with the tuner, and about every note, starting from the second, it is written down - whether it is closer or farther to the triangle sound than the previous note. It is known in advance that the frequency of Maxima triangle sound is not less than 30 hertz and not more than 4000 hertz.

You need to write a program that determines in what interval the frequency of the sound of the triangle can be.

Input format
The first line of the input file contains an integer n - the number of notes played by Maxim with the tuner (2 ≤ n ≤ 1000).
The following n lines contain Maxim's recordings. Each line has two components: real number fi - the frequency in Hertz set on the tuner (30 ≤ fi ≤ 4000), and the word "closer" or the word "further" for each frequency except the first.

The word "closer" means that the frequency of a given note is closer to the frequency of the triangle than the frequency of the previous note, which is described formally by the relation: |fi - ftriangle| < |fi - 1 - ftriangle| .

The word "further" means that the frequency of this note is further than the previous one.

If it turns out that the next note is as close to the triangle sound as the previous note, Maxim could write down any of the two words above.

It is guaranteed, that the results obtained by Maxim are consistent.

Output Format
The output file must contain two real numbers, the smallest and the largest possible value of the frequency of the sound of the triangle produced by Maxim. The numbers should be printed with accuracy not worse than 10^6.


Input: 

3  // number of notes

440


220 closer - 30
300 further


1.  30..................440................4000
2.  30..!.220...........440................4000
3.  30..!.220...300.....440................4000
           |     |
         <-c     f->

excepted: ??
|note_played|status|mid_point|
|:----|:----|:----|
|440|-||
|220|c|< 330|220-330 < 0|
|300|f|< 260|300-260 > 0|

low=30, high=4000

low=min(low, f(mp)), highest=max(4000, f(mp))

|note_played|status|mid_point|comparator|
|:----|:----|:----|:----|
|554|-|||
|880|f|717|'lt'| 
|440|c|660|'lt'|
|622|c|531|'gt'| 

up_bound := 40_000 (right_pointer)
low_bound := 30 (left_pointer)

'if' check_case 'then' low_bound = max(low_bound,mid_point) 'else' up_bound=min(up_bound,mid_point)

check_cases
1. 'if' f & note_played > mid_point 'then' up_bound = min(up_bound,mid_point)
2. 'if' f & note_played < mid_point 'then' low_bound = max(low_bound,mid_point)
3. 'if' c & note_played > mid_point 'then' low_bound = max(low_bound,mid_point)
4. 'if' c & note_played < mid_point 'then' up_bound = min(up_bound,mid_point)