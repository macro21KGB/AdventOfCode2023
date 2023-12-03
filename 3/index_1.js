import fs from "fs/promises";

const inputData = await fs.readFile("./input.txt", "utf-8");

const possibleSymbols = {
    "*": true,
    "@": true,
    "#": true,
    "$": true,
    "+": true,
    "&": true,
    "%": true,
    "/": true,
    "-": true,
    "=": true
};

const lines = inputData.split("\n");
// const regexNumberWithSymbol = /\d+[*@#\*\$\+&%\/\-\=]/g;
// const regex = /(\d+)(?:\.+)?(?:\n)?[*@#\*\$\+&%\/\-\=]/g;
const digitRegex = /\d+/g;

/**
 * check if a symbol is adiacent to a number horizontally
 * @param {number} currentLineIndex 
 * @param {number} startIndex 
 * @param {number} endIndex 
 * @returns boolean
 */
const checkSymbolAdiacentHorizontally = (currentLineIndex, startIndex, endIndex) => {
    const leftSide = startIndex - 1 < 0 ? startIndex : startIndex - 1;
    const rightSide = endIndex + 1 > lines[currentLineIndex].length ? endIndex : endIndex + 1;
    const selectedNumber = lines[currentLineIndex].slice(leftSide, rightSide);
    const leftRightSymbol = /[@#\*\$\+&%\/\-\=]/g

    if (selectedNumber.match(leftRightSymbol)) {
        return true;
    }

    return false;
}

const checkSymbolAdiacentVertically = (currentLineIndex, startIndex, endIndex) => {
    const topSide = currentLineIndex - 1 < 0 ? null : currentLineIndex - 1;
    const bottomSide = currentLineIndex + 1 > lines.length - 1 ? null : currentLineIndex + 1;
    const topBottomSymbol = /[\@\#\*\$\+&%\/\-\=]/g

    const offsettedStartIndex = startIndex - 1 < 0 ? startIndex : startIndex - 1;
    const offsettedEndIndex = endIndex + 1 > lines[currentLineIndex].length ? endIndex : endIndex + 1;

    if (topSide && lines[topSide].slice(offsettedStartIndex, offsettedEndIndex).match(topBottomSymbol)) {
        return true;
    }

    if (bottomSide && lines[bottomSide].slice(offsettedStartIndex, offsettedEndIndex).match(topBottomSymbol)) {
        return true;
    }

    return false;

}

let totalSum = 0;
for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    const digitsFoundInLine = line.matchAll(digitRegex);
    if (digitsFoundInLine !== null) {
        [...digitsFoundInLine].forEach((match) => {
            const startIndex = match.index;
            const endIndex = startIndex + match[0].length;

            const isSymbolAdiacentHorizontally = checkSymbolAdiacentHorizontally(i, startIndex, endIndex);
            const isSymbolAdiacentVertically = checkSymbolAdiacentVertically(i, startIndex, endIndex);

            if (isSymbolAdiacentHorizontally) {
                totalSum += parseInt(match);
            }

            if (isSymbolAdiacentVertically) {
                totalSum += parseInt(match);
            }
        });
    }
}

console.log(totalSum);