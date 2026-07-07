package lasagnamaster

func PreparationTime(layers []string, layerPrepTime int) int {
    if layerPrepTime == 0 {
        layerPrepTime = 2
    }
    return len(layers) * layerPrepTime;
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        switch layer {
        case "sauce":
            sauce += 0.2
        case "noodles":
            noodles += 50
        }
    }
    return
}

func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1] 
}

func ScaleRecipe(quantities []float64, portions int)(scaledQuantities []float64){
    scaledQuantities = make([]float64, len(quantities))
    for i, q := range quantities {
        scaledQuantities [i] = (q / 2.0) * float64(portions)
    }
    return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
