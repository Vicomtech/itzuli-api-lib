import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="itzuli-pkg-vicomtech", # Replace with your own username
    version="1.0.1",
    author="Vicomtech.org",
    author_email="info@vicomtech.org",
    description="Itzuli API library",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/Vicomtech/itzuli-api-lib/",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: GPL3 License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',
)
