# tests/core-tests/test_cnm_core.py

import unittest
from core.cnm_core import CNMCore

class TestCNMCore(unittest.TestCase):
    def test_load_config(self):
        cnm_core = CNMCore()
        config = cnm_core.load_config()

        self.assertIsInstance(config, dict)
        self.assertIn('version', config)
        self.assertIn('name', config)

    def test_process_data(self):
        cnm_core = CNMCore()
        processed_data = cnm_core.process_data('test_data')

        self.assertEqual(processed_data, 'Processed: test_data')

if __name__ == '__main__':
    unittest.main()
