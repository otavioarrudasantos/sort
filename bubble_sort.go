package sort

func BubbleSort(data []int){
    for i := 0; i < len(data); i ++{
        for j := 0; j < len(data) -1; j++{
            if data[j] > data[j+1]{
                var aux = data[j]
                data[j] = data[j+1]
                data[j+1] = aux;
            }
        }
    }
}
