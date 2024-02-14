// frontend/cnm/modules/feature1.js

import { formatDate } from '../utils.js';

function feature1Module() {
  console.log('Feature 1 Module Loaded');

  const currentDate = new Date();
  const formattedDate = formatDate(currentDate);

  console.log('Feature 1 Module: Current Date is', formattedDate);
}

export { feature1Module };
