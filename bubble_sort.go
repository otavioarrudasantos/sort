package sort

func BubbleSort(data []int){
    length := len(data);
    for i := 0; i < length; i ++{
        for j := 0; j < length -1; j++{
            if data[j] > data[j+1]{
                var aux = data[j]
                data[j] = data[j+1]
                data[j+1] = aux
            }
        }
    }
}
