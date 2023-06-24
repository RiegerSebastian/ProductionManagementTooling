module example.com/main

go 1.20

require example.com/productionCalculation v0.0.0-00010101000000-000000000000

replace example.com/productionCalculation => ../productionCalculation

replace example.com/equipmentData => ../equipmentData

require example.com/equipmentData v0.0.0-00010101000000-000000000000
