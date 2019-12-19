echo start 

echo LoginServer:登录服务器启动(http)
start ZhjGo.exe 8891 DT

echo GateWay:网关服务器启动(websocket)
start ZhjGo.exe 8888 GW

echo DBServer:数据库服务器启动(rpc)
start ZhjGo.exe 8890 DB

echo Global Server:公共服务器启动(websocket，内服务)
start ZhjGo.exe 8894 GL

echo GM server :服务器启动(http)
start ZhjGo.exe 8892 GM

echo DSQ server :服务器启动(websocket)
echo ZhjGo.exe 8895 DSQ

exit