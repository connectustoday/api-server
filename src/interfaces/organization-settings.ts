import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';
import IAccountSettings from './account-settings';

export default class IOrganizationSettings extends IAccountSettings {

	@prop({ required: true })
	is_nonprofit: boolean;

}
