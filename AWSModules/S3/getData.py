import os
import platform
import subprocess
import time

def get_os_type():
    return platform.system().lower()

def get_os_version():
    return platform.release()

def get_cpu_info():
    if get_os_type() == 'windows':
        return subprocess.check_output("wmic cpu get caption", shell=True).decode().strip()
    elif get_os_type() == 'linux':
        return subprocess.check_output("lscpu | grep 'Model name'", shell=True).decode().strip()
    elif get_os_type() == 'darwin':
        return subprocess.check_output("sysctl -n machdep.cpu.brand_string", shell=True).decode().strip()
    else:
        return "Unknown CPU Info"
    
def get_cpu_type():
    return platform.processor()

def get_primary_interface():
    if get_os_type() == 'windows':
        return subprocess.check_output("wmic nic where (NetEnabled=true) get Name", shell=True).decode().strip()
    elif get_os_type() == 'linux':
        return subprocess.check_output("ip route | grep default | awk '{print $5}'", shell=True).decode().strip()
    elif get_os_type() == 'darwin':
        return subprocess.check_output("route -n get default | grep interface | awk '{print $2}'", shell=True).decode().strip()
    else:
        return "Unknown Primary Interface"
    

def get_system_info():
    system_info = {
        "os_type": get_os_type(),
        "os_version": get_os_version(),
        "cpu_info": get_cpu_info(),
        "cpu_type": get_cpu_type(),
        "primary_interface": get_primary_interface(),
        "timestamp": time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    }
    return system_info

print("Gathering system information...")
system_info = get_system_info()
print("System Information:")
for key, value in system_info.items():
    print(f"{key}: {value}")
# This script gathers system information such as OS type, version, CPU info, and primary network interface.
# It uses platform and subprocess modules to retrieve the necessary details.
# The information is printed in a structured format.
# The script is designed to run on Windows, Linux, and macOS.
# Note: Ensure you have the necessary permissions to run system commands.
# The script can be extended to include more system details as needed.
# The script is intended for educational purposes and may require adjustments based on specific system configurations.
# The script is designed to be run in a Python environment with access to system commands.
# The script can be saved as getdata.py and executed in a terminal or command prompt.
# Ensure that the script is run with appropriate permissions to access system information.
# The script can be modified to log the output to a file or send it to a remote server for further analysis.
# The script can be used for system diagnostics, monitoring, or inventory purposes.
