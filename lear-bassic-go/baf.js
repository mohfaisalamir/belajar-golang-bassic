function ganjil_not58(length){
    let i=0
    while(i<=length){
        if(i%2!=0 && i%5==0 && i%8==0){
            console.log(i)
        }
        i++
    }

}
ganjil_not58(100)