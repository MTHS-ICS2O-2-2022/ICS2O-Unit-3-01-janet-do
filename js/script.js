// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file generates the area of a trapezoid

"use strict"

function calculate() {
  // input
  const sideA = parseInt(document.getElementById("sideA").value)
  const sideB = parseInt(document.getElementById("sideB").value)
  const height = parseInt(document.getElementById("height").value)

  // process
  const area = ((sideA * sideB) / 2) * height

  // output
  document.getElementById("area").innerHTML = "Area is: " + area + " cm²"
}
