# mean_response_tme
**Enhancing Cluster Performance by Reducing Response Time Variance Issues**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** Need to add conf name*****************
- **Publication Date:** Feb, 2026
- **ISSN:**  2147-6799

### Abstract
Distributed cluster systems often experience unstable performance due to response time variance caused by synchronization delays, cross node communication, and uneven workload distribution. Existing static sharding approaches increase coordination overhead as cluster size grows, leading to rising response times and unpredictable latency. This study introduces a variance aware framework combining adaptive shard management, locality aware scheduling, and variance focused monitoring. Experimental evaluation across clusters of 3 to 11 nodes shows improved response time stability and more predictable system performance in distributed environments.

### Core Innovations Presented in the Study
- **Variance Aware Performance Framework:**  
Introduced a system level framework that focuses on reducing response time variance by combining adaptive shard placement, locality conscious scheduling, and variance oriented telemetry analysis.

- **Locality Driven Transaction Scheduling:**  
Developed a scheduling strategy that prioritizes transactions based on node load and data locality to minimize cross node coordination and synchronization delays.

- **Prototype Based Cluster Evaluation:** 
Implemented a distributed cluster simulation using Go based processing to evaluate response time behavior and workload distribution across different cluster configurations.

- **Scalability Assessment Across Node Configurations:**  
Conducted experiments on clusters containing 3, 5, 7, 9, and 11 nodes to examine how adaptive scheduling stabilizes response time as cluster size increases.

### System Level Value and Operational Benefits
- **Lower Mean Response Time:**
The adaptive scheduling framework significantly reduces mean response time compared with static sharding architectures in distributed clusters.

- **Improved Performance Predictability:**  
Variance focused monitoring and adaptive scheduling reduce jitter, enabling more stable and consistent response times across workloads.

- **Enhanced Scalability in Cluster Environments:**  
Response time growth remains controlled as cluster size increases, demonstrating that variance aware mechanisms support reliable scaling.

- **Applicability to Mission Critical Distributed Platforms:**  
The framework is suitable for distributed databases, financial systems, cloud infrastructures, and real time analytics platforms that require predictable performance and stable response times.

### Experimental Results (Summary)

  | Nodes | Lock Based 2PL CPU %| Lightweight Runtime Detection %| Improvment (%) |
  |-------|---------------------| -------------------------------| ---------------|
  | 3     |  88                 | 55                             | 37.50          |
  | 5     |  84                 | 49                             | 41.67          |
  | 7     |  82                 | 46                             | 43.90          |
  | 9     |  80                 | 43                             | 46.25          |
  | 11    |  79                 | 41                             | 48.10          |

### Citation
Lightweight Runtime Conflict Detection for CPU Efficient Transaction Processing
* Naveen Kumar Bandaru
* International Journal of Intelligent Systems and Applications in Engineering (IJISAE) 
* ISSN 2147-6799
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijisae.org/index.php/IJISAE \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com







