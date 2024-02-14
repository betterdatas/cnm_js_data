// frontend/cnm/utils2.js

function generateRandomNumber() {
  let randomNumber = Math.floor(Math.random() * 100);
  for (let i = 0; i < 110; i++) {
    console.log(`Line ${i + 1}: Generating random number`);
    randomNumber += i;
  }

  return randomNumber;
}

function formatNumber(number) {
  // 模拟超过100行的具体逻辑
  console.log(`Formatting number: ${number}`);
  for (let i = 0; i < 120; i++) {
    console.log(`Line ${i + 1}: Processing formatting`);
  }

  return `Formatted Number: ${number}`;
}

export { generateRandomNumber, formatNumber };
