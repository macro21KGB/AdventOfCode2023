import fs from "fs/promises";

/**
 * @typedef {{start: number, end: number, number: number, lineIndex: number}} Part
 */

/*
A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.
This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.
What is the sum of all of the gear ratios in your engine schematic?
*/

/**
 * 
 * @param {number} lineIndex
 * @param {number} startIndex 
 * @param {number} endIndex 
 * @param {Map<number, Part[]} mapOfParts 
 * @returns {Part | null}
 */
const checkPartsLine = (lineIndex, startIndex, endIndex, mapOfParts) => {
    const leftStart = startIndex - 1 < 0 ? 0 : startIndex - 1;
    const rightEnd = endIndex + 1 > splittedInputData.length - 1 ? endIndex : endIndex + 1;
    const stringToSearchSplitted = splittedInputData[lineIndex].slice(leftStart, rightEnd).split("");

    const parts = mapOfParts.get(lineIndex);
    let foundPart = null;

    if (!parts)
        return null;

    stringToSearchSplitted.forEach((char, index) => {
        if (char !== "*" && char !== ".") {
            if (checkIfPositionInteresectPart(lineIndex, leftStart + index, mapOfParts)) {
                const part = parts.find((part) => {
                    return index >= part.start && index <= part.end;
                });

                if (part) {
                    foundPart = part;
                    return;
                }
            }
        }

    });

    return foundPart;
};

const inputData = await fs.readFile("./test_input1.txt", "utf-8");
const splittedInputData = inputData.split("\n");


const allDigitsMatchList = inputData.matchAll(/\d+/g);

const allDigits = [...allDigitsMatchList];


const partsMap = new Map();
allDigits.forEach((digit) => {
    const currentLineIndex = inputData.slice(0, digit.index).split("\n").length - 1;
    if (partsMap.has(currentLineIndex)) {
        partsMap.get(currentLineIndex).push({
            start: digit.index,
            end: digit.index + digit[0].length,
            number: parseInt(digit[0]),
            lineIndex: currentLineIndex
        });
    } else {
        partsMap.set(currentLineIndex, [{
            start: digit.index,
            end: digit.index + digit[0].length,
            number: parseInt(digit[0]),
            lineIndex: currentLineIndex
        }]);
    }
});


/**
 * 
 * @param {number} lineIndex 
 * @param {number} selectedIndex 
 * @param {Map<number, Part[]>} mapOfParts 
 */
const checkIfPositionInteresectPart = (lineIndex, selectedIndex, mapOfParts) => {

    const parts = mapOfParts.get(lineIndex);
    const lengthOfLine = splittedInputData[lineIndex].length;

    if (!parts)
        return false;

    const part = parts.find((part) => {
        return selectedIndex >= part.start % lengthOfLine && selectedIndex <= part.end % lengthOfLine;
    });

    return !!part;
};

inputData.split("\n").forEach((line, indexLine) => {
    const matches = line.matchAll(/\*/g);

    const gearSymbols = [...matches];

    gearSymbols.forEach((gearSymbol) => {
        const startIndex = gearSymbol.index;
        const endIndex = gearSymbol.index + 1;

        // check every parts around the gear, in total of 26 spaces
        let currentLineIndex = indexLine;


        const upLine = currentLineIndex - 1 < 0 ? 0 : currentLineIndex - 1;
        const downLine = currentLineIndex + 1 > splittedInputData.length - 1 ? currentLineIndex : currentLineIndex + 1;

        const partUp = checkPartsLine(upLine, startIndex, endIndex, partsMap);
        const partDown = checkPartsLine(downLine, startIndex, endIndex, partsMap);
        const partCurrent = checkPartsLine(currentLineIndex, startIndex, endIndex, partsMap);

        // if excatly 2 parts around the gear, then it's a gear and we can calculate the ratio
        const partsAroundGear = [partUp, partDown, partCurrent].filter((part) => part !== null);

        if (partsAroundGear.length === 2) {
            const gearRatio = partsAroundGear[0].number * partsAroundGear[1].number;
        }


    });


});


