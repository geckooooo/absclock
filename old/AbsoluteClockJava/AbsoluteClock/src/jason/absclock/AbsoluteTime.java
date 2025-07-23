package jason.absclock;

import java.util.*;

public class AbsoluteTime {
	private Date systemTime;
	
	public Date getSystemTime() {
		return systemTime;
	}

	public void setSystemTime(Date systemTime) {
		this.systemTime = systemTime;
	}

	/**
	 * Returns eternity: "E" for the duration of our universe.
	 * 
	 * @return the eternity value.
	 */
	public String getEternity() {
		return "E";
	}

	/**
	 * Returns the current epoch, which is treated as constant for the life
	 * of this program.
	 * 
	 * @return the epoch value.
	 */
	public String getEpoch() {
		return "0";
	}

	public String getYear() {
		return systemTime.toString();
	}
	
	/**
	 * Override toString as a convenience. This returns the entire
	 * string represenation of absolute time.
	 */
	public String toString() {
		return getEternity() + ":"
			+ getEpoch() + ":"
			+ getYear();
	}
	
	public AbsoluteTime() {
		systemTime = new Date();
	}
}