/**
 * Created by nathan.shumoogum on 19/06/2017.
 */
import org.apache.ignite.Ignite;
import org.apache.ignite.IgniteCache;
import org.apache.ignite.IgniteCompute;
import org.apache.ignite.IgniteException;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

import static org.apache.ignite.Ignition.start;

public class loadApacheIgnite {
    public static void main(String[] args) throws IgniteException {
        try (Ignite ignite = start("examples/config/example-ignite.xml")) {
            // Put values in cache.
            IgniteCache<Integer, String> cache = ignite.getOrCreateCache("myCache");
            IgniteCompute compute = ignite.compute(ignite.cluster().forRemotes());

            String csvFile = "/Users/nathan.shumoogum/ons/go/src/github.com/ONSdigital/dp-cmd-database-spikes/input-files/ASHE07E_2013WARDH_2015_3_EN_Earnings_just_Statistics.csv";

            BufferedReader br = null;
            String line = "";
            String cvsSplitBy = ",";
            String[] csvHeaders = new String[21];
            Map<String, String> data = new HashMap<String, String>();

            try {

                br = new BufferedReader(new FileReader(csvFile));
                Integer count = 0;
                while ((line = br.readLine()) != null) {
                    // use comma as separator
                    String[] value = line.split(cvsSplitBy);

                    if (count == 0) {
                        for (int i = 0; i < value.length; i++) {
                            // Add values to hash
                            csvHeaders[i] = value[i];
                            System.out.println("Header is: " + value[i]);
                        }
                        count++;
                    } else {
                        String concatenatedDataSet = "{";
                        for (int i = 0; i < value.length; i++) {
                            data.put(csvHeaders[i], value[i]);
                            //System.out.println("------- \n value of observation is: " + data.get(csvHeaders[0]));
                            concatenatedDataSet = concatenatedDataSet + "\"" + csvHeaders[i] + "\": \"" + value[i] + "\",";

                        }

                        concatenatedDataSet = concatenatedDataSet.substring(0, concatenatedDataSet.length() - 1) + "}";

                        //System.out.println("------- \n concatenated data set is: " + concatenatedDataSet);

                        //cache.put(count, doc);
                        cache.put(count, concatenatedDataSet);
                        // Get values from cache
                        // Broadcast 'Hello World' on all the nodes in the cluster.
                        Integer finalCount = count;
                        //System.out.println("------- \n id is: " + finalCount);
                        compute.broadcast(() -> System.out.println(cache.get(finalCount)));

                        count++;
                    }
                }

                System.out.println("------- \n id is: " + count);

            } catch (FileNotFoundException e) {
                e.printStackTrace();
            } catch (IOException e) {
                e.printStackTrace();
            } finally {
                if (br != null) {
                    try {
                        br.close();
                    } catch (IOException e) {
                        e.printStackTrace();
                    }
                }
            }
        }
    }
}


