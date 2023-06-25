const temp = document.getElementById("animate");
let initialContent = temp.outerHTML
console.log(initialContent)

function myMove() {
    let id = null;
    const elem = document.getElementById("animate");
    elem.outerHTML = initialContent
    console.log(elem)
    let posX = 0;
    let posY = 0;
    let stop = 0;
    clearInterval(id);
    console.log(initialContent)
    id = setInterval(frame, 1);
    function frame() {
        if (stop == 0) {
            if (posX == 350 && posY == 350) {
                elem.style.backgroundColor = "blue";
                stop = 1;
            } else {
                stop0()
            }

        } else if (stop == 1) {
            if (posX == 350 && posY == 0) {
                elem.style.backgroundColor = "green";
                stop = 2
            } else {
                stop1()
            }

        } else if (stop == 2) {
            if (posX == 0 && posY == 350) {
                elem.style.backgroundColor = "yellow";
                stop = 3
            } else {
                stop2()
            }

        } else if (stop == 3) {
            if (posX == 0 && posY == 0) {
                elem.style.backgroundColor = "red";
                stop = 0
            } else {
                stop3()
            }
        }
    }
    function stop0() {
        posX++;
        posY++;
        elem.style.top = posY + "px";
        elem.style.left = posX + "px";
    }
    function stop1() {
        posY--;
        elem.style.top = posY + "px";
        elem.style.left = posX + "px";
    }
    function stop2() {
        posX--;
        posY++;
        elem.style.top = posY + "px";
        elem.style.left = posX + "px";
    }
    function stop3() {
        posY--;
        elem.style.top = posY + "px";
        elem.style.left = posX + "px";
    }
}