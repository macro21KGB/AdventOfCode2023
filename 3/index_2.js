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
 * @returns {Part[] | null}
 */
const checkPartsLine = (lineIndex, startIndex, endIndex, mapOfParts) => {
    const parts = mapOfParts.get(lineIndex);
    let foundParts = [];

    if (!parts)
        return null;


    parts.forEach((part) => {
        if (part.end >= startIndex - 1 && part.start <= endIndex) {
            foundParts.push(part);
        }
    });

    return foundParts ? foundParts : null;
};




/**
 * 
 * @param {string} inputData 
 * @returns {Map<number, Part[]>}
 */
const createPartsMap = (inputData) => {
    const splittedInputData = inputData.split("\n");
    const partsMap = new Map();

    splittedInputData.forEach((line, indexLine) => {
        const matches = line.matchAll(/\d+/g);

        const digits = [...matches];

        digits.forEach((digit) => {
            const startIndex = digit.index;
            const endIndex = digit.index + digit[0].length - 1;
            const part = {
                start: startIndex,
                end: endIndex,
                number: parseInt(digit[0]),
                lineIndex: indexLine
            };

            if (partsMap.has(indexLine)) {
                const parts = partsMap.get(indexLine);
                parts.push(part);
                partsMap.set(indexLine, parts);
            } else {
                partsMap.set(indexLine, [part]);
            }
        });
    });

    return partsMap;
}

// iterate over every line

/**
 * 
 * @param {string} inputData 
 */
const checkGear = (inputData) => {
    const allGearsValid = [];
    const splittedInputData = inputData.split("\n");
    const partsMap = createPartsMap(inputData);
    let totalSumOfGearRatio = 0;

    inputData.split("\n").forEach((line, indexLine) => {
        const matches = line.matchAll(/\*/g);

        const gearSymbols = [...matches];

        gearSymbols.forEach((gearSymbol) => {
            const startIndex = gearSymbol.index;
            const endIndex = gearSymbol.index + 1;
            const partsAroundGear = []

            // check every parts around the gear, in total of 26 spaces
            let currentLineIndex = indexLine;


            const upLine = currentLineIndex - 1 < 0 ? null : currentLineIndex - 1;
            const downLine = currentLineIndex + 1 > splittedInputData.length - 1 ? null : currentLineIndex + 1;

            if (upLine !== null) {
                const partUp = checkPartsLine(upLine, startIndex, endIndex, partsMap);
                if (partUp) {
                    partUp.forEach((part) => {
                        partsAroundGear.push(part.number);
                    });

                }
            }

            if (downLine !== null) {
                const partDown = checkPartsLine(downLine, startIndex, endIndex, partsMap);
                if (partDown) {
                    partDown.forEach((part) => {
                        partsAroundGear.push(part.number);
                    });
                }
            }

            const partCurrent = checkPartsLine(currentLineIndex, startIndex, endIndex, partsMap);
            if (partCurrent) {
                partCurrent.forEach((part) => {
                    partsAroundGear.push(part.number);
                });
            }
            // if excatly 2 parts around the gear, then it's a gear and we can calculate the ratio
            if (partsAroundGear.length === 2) {
                const gearRatio = partsAroundGear[0] * partsAroundGear[1];
                allGearsValid.push(partsAroundGear)
                totalSumOfGearRatio += gearRatio;
            }


        });


    });


    return {
        allGearsValid,
        totalSumOfGearRatio
    };
}


const inputDataRaw = await fs.readFile("./input.txt", "utf-8");


const {
    allGearsValid,
    totalSumOfGearRatio
} = checkGear(inputDataRaw);

console.log(totalSumOfGearRatio);

export {
    checkPartsLine,
    checkGear
}