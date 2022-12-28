### Cplus Test

#### 单元测试框架Google Test
- 下载  
  去https://github.com/google/googletest/releases下载，千万不要直接git clone。可以下载v1.10.0版本： https://github.com/google/googletest/archive/refs/tags/release-1.10.0.zip
- 安装  
  ```
  cd googletest-release-1.10.0 
  cmake .
  make
  make install 
  ```
#### 编写单元测试脚本

#### 编译单元测试脚本
- 使用CMakeLists.txt  
- 使用命令编译
```
g++ -Wall -std=c++11 test.cpp -lgtest -lgtest_main -lpthread -o test
```
