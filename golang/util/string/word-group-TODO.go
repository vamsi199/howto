// given a slice of single words, form groups on n subsequent words. for example: 
// input example
// input slice is ["Hello", "Gopher", "How", "are", "you", "doing"]
// output examples
// output for n = 1 is same as input [["Hello"], ["Gopher"], ["How"], ["are"], ["you"], ["doing"]]
// output for n = 2 is [["Hello" "Gopher"], ["Gopher" "How"], ["How" "are"], ["are" "you"], ["you" "doing"]]
// output for n = 3 is [["Hello" "Gopher" "How"], ["Gopher" "How" "are"], ["How" "are" "you"], ["are" "you" "doing"]]
// signture:
// function sugnature is: WordGroup([]string, int)[][]string
