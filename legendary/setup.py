from setuptools import setup
from Cython.Build import cythonize

setup(
    ext_modules=cythonize("app.py")  # Replace 'app.py' with your Python file(s)
)
