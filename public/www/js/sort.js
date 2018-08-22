var count = 500;
var arr = [];
var pr,h,w,r;
function setup() {
  createCanvas(720, 400);
  for(var x=0; x < count; x++){
    arr.push(x);
  }
  background(0);   // Set the background to black
  r = shuffle(arr);
  updateFull(arr)
}
function updateFull(nr) {
  clear();
  background(0); 
  for(var x=0; x < nr.length; x++){
    pr = arr[x] * 100 / r.length;
    npr = nr[x] * 100 / r.length;
    h = height * npr / 100;
    w = width * pr / 100;
    c = 255 * npr / 100;
    stroke(c,200,0);
    strokeWeight(10);
    line(w-1, height+6, w+1, height-h+6);
  }
}
function updateLine(nr,i) {
  pr = arr[i] * 100 / r.length;
  npr = nr[i] * 100 / r.length;
  h = height * npr / 100;
  w = width * pr / 100;
  c = 255 * npr / 100;
  stroke(c,200,0);
  strokeWeight(10);
  line(w-1, height+6, w+1, height-h+6);
}
function draw() {
  background(0);   // Set the background to black
  /* Быстрая сортировка */
  // mergeSort (r)
  /* Пузырьковая сортировка */
  // for(var x=0; x < r.length; x++){
  //   if (r[x] > r[x + 1]) {
  //     var a = r[x]
  //     var b = r[x + 1]
  //     r[x] = b
  //     r[x + 1] = a
  //   }
  //   updateLine(r,x);
  // }


  /* Коктель шейкера */
  for(var i = 0; i < r.length - 2; i++) {
    if(r[i] > r[i+1]) {
      var temp = r[i];
      r[i] = r[i+1];
      r[i+1] = temp;
      swapped = true;
    }
  }	
  if(!swapped) {
    r = shuffle(arr);
    updateFull(r)
  }
  swapped = false;
  for( i = r.length - 2; i > 0; i--) {
    if(r[i] > r[i+1]) {
      var temp1 = r[i];
      r[i] = r[i+1];
      r[i+1] = temp1;
      swapped = true;
    }
  }
   
  updateFull(r)
}

function shuffle(array) {
  var currentIndex = array.length, temporaryValue, randomIndex;

  // While there remain elements to shuffle...
  while (0 !== currentIndex) {

    // Pick a remaining element...
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex -= 1;

    // And swap it with the current element.
    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }

  return array;
}
function mergeSort (arr) {
  if (arr.length === 1) {
    // return once we hit an array with a single item
    return arr
  }

  const middle = Math.floor(arr.length / 2) // get the middle item of the array rounded down
  const left = arr.slice(0, middle) // items on the left side
  const right = arr.slice(middle) // items on the right side

  return merge(
    mergeSort(left),
    mergeSort(right)
  )
}

// compare the arrays item by item and return the concatenated result
function merge (left, right) {
  let result = []
  let indexLeft = 0
  let indexRight = 0

  while (indexLeft < left.length && indexRight < right.length) {
    if (left[indexLeft] < right[indexRight]) {
      result.push(left[indexLeft])
      indexLeft++
      updateLine(result,indexLeft)
    } else {
      result.push(right[indexRight])
      indexRight++
      updateLine(result,indexRight)
    }
  }

  return result.concat(left.slice(indexLeft)).concat(right.slice(indexRight))
}